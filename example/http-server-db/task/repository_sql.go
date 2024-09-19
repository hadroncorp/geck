package task

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	gecksql "github.com/hadroncorp/geck/data/sql"
	"github.com/hadroncorp/geck/systemerror"
)

type RepositorySQL struct {
	Client gecksql.Client
}

var _ Repository = (*RepositorySQL)(nil)

func NewRepositorySQL(client gecksql.Client) RepositorySQL {
	return RepositorySQL{
		Client: client,
	}
}

func (r RepositorySQL) Save(ctx context.Context, entity Task) error {
	var stmt string
	var args []any
	if entity.GetVersion() > 0 {
		stmt = "UPDATE tasks SET task_name=$1,status=$2,version=$3 WHERE task_id=$4"
		args = []any{entity.Name, entity.Status, entity.Version, entity.ID}
	} else {
		stmt = "INSERT INTO tasks(task_id,task_name,status) VALUES ($1,$2,$3)"
		args = []any{entity.ID, entity.Name, entity.Status}
	}
	_, err := r.Client.ExecContext(ctx, stmt, args...)
	if err != nil && strings.HasPrefix(err.Error(), "ERROR: duplicate key") {
		return systemerror.NewResourceAlreadyExists[Task](entity.ID)
	}
	return err
}

func (r RepositorySQL) SaveMany(ctx context.Context, entities []Task) error {
	errs := make([]error, len(entities))
	for _, entity := range entities {
		if err := r.Save(ctx, entity); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (r RepositorySQL) Remove(ctx context.Context, entity Task) error {
	stmt := "DELETE FROM tasks WHERE task_id=$1"
	_, err := r.Client.ExecContext(ctx, stmt, entity.ID)
	return err
}

func (r RepositorySQL) FindByKey(ctx context.Context, key string) (*Task, error) {
	stmt := "SELECT task_id,task_name,status,create_time,create_by,last_update_time,last_update_by,is_active,version FROM tasks WHERE task_id=$1"
	row := r.Client.QueryRowContext(ctx, stmt, key)
	ent := &Task{}
	err := row.Scan(&ent.ID, &ent.Name, &ent.Status, &ent.CreateTime, &ent.CreateBy, &ent.LastUpdateTime, &ent.LastUpdateBy, &ent.IsActive, &ent.Version)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return ent, nil
}

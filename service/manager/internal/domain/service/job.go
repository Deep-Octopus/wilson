package service

import (
	"github.com/dstgo/wilson/framework/kratosx"

	"github.com/dstgo/wilson/api/gen/errors"
	"github.com/dstgo/wilson/service/manager/internal/conf"
	"github.com/dstgo/wilson/service/manager/internal/domain/entity"
	"github.com/dstgo/wilson/service/manager/internal/domain/repository"
	"github.com/dstgo/wilson/service/manager/internal/types"
)

type Job struct {
	conf *conf.Config
	repo repository.Job
}

func NewJob(config *conf.Config, repo repository.Job) *Job {
	return &Job{conf: config, repo: repo}
}

// ListJob 获取职位信息列表
func (u *Job) ListJob(ctx kratosx.Context, req *types.ListJobRequest) ([]*entity.Job, uint32, error) {
	list, total, err := u.repo.ListJob(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "list job error", "err", err.Error())
		return nil, 0, errors.ListError()
	}
	return list, total, nil
}

// CreateJob 创建职位信息
func (u *Job) CreateJob(ctx kratosx.Context, req *entity.Job) (uint32, error) {
	id, err := u.repo.CreateJob(ctx, req)
	if err != nil {
		ctx.Logger().Warnw("msg", "create job error", "err", err.Error())
		return 0, errors.CreateErrorWrap(err)
	}
	return id, nil
}

// UpdateJob 更新职位信息
func (u *Job) UpdateJob(ctx kratosx.Context, req *entity.Job) error {
	if err := u.repo.UpdateJob(ctx, req); err != nil {
		ctx.Logger().Warnw("msg", "update job error", "err", err.Error())
		return errors.UpdateErrorWrap(err)
	}
	return nil
}

// DeleteJob 删除职位信息
func (u *Job) DeleteJob(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteJob(ctx, id); err != nil {
		ctx.Logger().Warnw("msg", "delete job error", "err", err.Error())
		return errors.DeleteErrorWrap(err)
	}
	return nil
}

package model

import "context"

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (int64, error) {
	res, err := m.conn.ExecCtx(ctx,
		"INSERT INTO users (account, password, created_at) VALUES (?, ?, NOW())",
		data.Account, data.Password,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	var u Users
	err := m.conn.QueryRowCtx(ctx,
		&u,
		"SELECT id, account, password, created_at FROM users WHERE id = ? LIMIT 1",
		id,
	)
	return &u, err
}

func (m *defaultUsersModel) Update(ctx context.Context, data *Users) error {
	_, err := m.conn.ExecCtx(ctx,
		"UPDATE users SET account=?, password=? WHERE id=?",
		data.Account, data.Password, data.Id,
	)
	return err
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	_, err := m.conn.ExecCtx(ctx, "DELETE FROM users WHERE id=?", id)
	return err
}

func (m *defaultUsersModel) List(ctx context.Context) ([]Users, error) {
	var resp []Users
	err := m.conn.QueryRowsCtx(ctx, &resp,
		"SELECT id, account, password, created_at FROM users ORDER BY id DESC")
	return resp, err
}

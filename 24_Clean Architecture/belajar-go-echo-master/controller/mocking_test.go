package controller_test

type UserService struct {
	db *gorm.DB
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	users := make([]model.User, 0)
	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func TestGetAllUsers(t *testing.T) {
	// create mock DB
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %s", err)
	}
	defer mockDB.Close()

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mockDB,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error creating mock DB connection: %s", err)
	}

	// create mock service
	service := &UserService{db}

	// expectations
	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(1, "John Doe", "john@example.com").
		AddRow(2, "Jane Doe", "jane@example.com")
	mock.ExpectQuery("SELECT \\* FROM `users`").WillReturnRows(rows)

	// test GetAllUsers
	users, err := service.GetAllUsers()
	if assert.NoError(t, err) {
		assert.Len(t, users, 2)

		// check user data
		assert.Equal(t, uint(1), users[0].ID)
		assert.Equal(t, "John Doe", users[0].Name)
		assert.Equal(t, "john@example.com", users[0].Email)

		assert.Equal(t, uint(2), users[1].ID)
		assert.Equal(t, "Jane Doe", users[1].Name)
		assert.Equal(t, "jane@example.com", users[1].Email)
	}
}

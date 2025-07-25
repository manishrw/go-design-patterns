package builder

type User struct {
	Name  string
	Email string
	Age   int
}

type UserBuilder struct {
	user User
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{user: User{}}
}

func (b *UserBuilder) Name(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) Email(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) Age(age int) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) Build() User {
	return b.user
}
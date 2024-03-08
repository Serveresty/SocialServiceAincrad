import React, { useState } from 'react';

const RegistrationComponent = () => {
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [sex, setSex] = useState('');
  const [email, setEmail] = useState('');
  const [phone, setPhone] = useState('');
  const [password, setPassword] = useState('');

  const handleRegistration = async () => {
    try {
      // Здесь можно добавить валидацию полей формы перед отправкой на бэкенд

      const response = await fetch('http://localhost:8080/registration', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          first_name: firstName,
          last_name: lastName,
          sex: sex,
          email: email,
          phone: phone,
          password: password,
        }),
      });

      if (response.ok) {
        // Успешная регистрация, можно обработать ответ, например, сохранить токен
        const data = await response.json();
        console.log('Успешная регистрация:', data);

        // Дополнительные действия после успешной регистрации
      } else {
        // Обработка ошибок
        console.error('Ошибка регистрации:', response.statusText);
      }
    } catch (error) {
      console.error('Ошибка при отправке запроса:', error);
    }
  };

  return (
    <div>
      <h2>Форма регистрации</h2>
      <form>
        <label>
          Имя:
          <input type="text" name="first_name" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
        </label>
        <br />
        <label>
          Фамилия:
          <input type="text" name="last_name" value={lastName} onChange={(e) => setLastName(e.target.value)} />
        </label>
        <br />
        <label>
          Пол:
          <input type="text" name="sex" value={sex} onChange={(e) => setSex(e.target.value)} />
        </label>
        <br />
        <label>
          Email:
          <input type="email" name="email" value={email} onChange={(e) => setEmail(e.target.value)} />
        </label>
        <br />
        <label>
          Телефон:
          <input type="tel" name="phone" value={phone} onChange={(e) => setPhone(e.target.value)} />
        </label>
        <br />
        <label>
          Пароль:
          <input
            type="password"
            name="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </label>
        <br />
        <button type="button" onClick={handleRegistration}>
          Зарегистрироваться
        </button>
      </form>
    </div>
  );
};

export default RegistrationComponent;
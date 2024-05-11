import React, { useState, useEffect, useContext } from 'react';
import { Navigate, useNavigate } from 'react-router-dom';
import Cookies from 'universal-cookie';
import { Filecontext } from '../contexts/Filecontext';

const AuthComponent = () => {
  const cookies = new Cookies();
  const history = useNavigate();
  
  useEffect(() => {
    const authToken = cookies.get('authToken');
    if (authToken) {
      history('/');
    }
  }, [cookies, history]);

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { setLogID } = useContext(Filecontext)

  const handleLogin = async () => {
    try {
      // Здесь можно добавить валидацию email и password перед отправкой на бэкенд

      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      if (response.ok) {
        const data = await response.json();
        console.log('Успешная авторизация:', data.token);
        setLogID(data.id)

        cookies.set('authToken', data.token, { path: '/'});
        history('/');
      } else {
        // Обработка ошибок
        console.error('Ошибка авторизации:', response.statusText);
      }
    } catch (error) {
      console.error('Ошибка при отправке запроса:', error);
    }
  };

  return (
    <div>
      <h2>Форма авторизации</h2>
      <form>
        <label>
          Email:
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </label>
        <br />
        <label>
          Пароль:
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </label>
        <br />
        <button type="button" onClick={handleLogin}>
          Войти
        </button>
      </form>
    </div>
  );
};

export default AuthComponent;
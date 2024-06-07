import React, { useState, useEffect, useContext } from 'react';
import { Link, Navigate, useNavigate } from 'react-router-dom';
import Cookies from 'universal-cookie';
import { Filecontext } from '../contexts/Filecontext';
import s from '../styles/auth.module.css'

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
    <div className={s.backgroundd}>
    <div className={s.wrapper}>
      <form>
        <h2>Login</h2>
          <div className={s.input_field}>
          <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
          <label>Enter your email</label>
        </div>
        <div className={s.input_field}>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
          <label>Enter your password</label>
        </div>
        <div className={s.forget}>
          <label>
            <input type="checkbox" id={s.remember} />
            <p>Remember me</p>
          </label>
          <a href="#">Forgot password?</a>
        </div>
        <button type="button" onClick={handleLogin}>Log In</button>
        <div className={s.register}>
          <p>Don't have an account? <Link to={`/sign-up`}>Register</Link></p>
        </div>
      </form>
    </div>
    </div>

    // <div>
    //   <h2>Авторизация</h2>
    //   <form>
    //     <label>
    //       Email:
    //       <input
    //         type="email"
    //         value={email}
    //         onChange={(e) => setEmail(e.target.value)}
    //       />
    //     </label>
    //     <br />
    //     <label>
    //       Пароль:
    //       <input
    //         type="password"
    //         value={password}
    //         onChange={(e) => setPassword(e.target.value)}
    //       />
    //     </label>
    //     <br />
    //     <button type="button" onClick={handleLogin}>
    //       Войти
    //     </button>
    //   </form>
    // </div>
  );
};

export default AuthComponent;
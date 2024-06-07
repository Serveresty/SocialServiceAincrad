import React, { useState, useEffect } from 'react';
import { Link, Navigate, useNavigate } from 'react-router-dom';
import Cookies from 'universal-cookie';
import s from '../styles/auth.module.css'

const RegistrationComponent = () => {
  const cookies = new Cookies();
  const history = useNavigate();
  
  useEffect(() => {
    const authToken = cookies.get('authToken');
    if (authToken) {
      history('/');
    }
  }, [cookies, history]);

  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [sex, setSex] = useState('');
  const [email, setEmail] = useState('');
  const [phone, setPhone] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');

  const handleRegistration = async () => {
    try {
      
      if(password != confirmPassword) {
        console.error('Ошибка регистрации:', "Passwords not equal");
        return;
      }

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
        const data = await response.json();
        console.log('Успешная регистрация:', data);
        history('/sign-in');
      } else {
        console.error('Ошибка регистрации:', response.statusText);
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
          <input type="text" value={firstName} onChange={(e) => setFirstName(e.target.value)} required/>
          <label>Enter your first name</label>
        </div>
        <div className={s.input_field}>
          <input type="text" value={lastName} onChange={(e) => setLastName(e.target.value)} required/>
          <label>Enter your last name</label>
        </div>
        <div className={s.input_field}>
          <input type="text" value={sex} onChange={(e) => setSex(e.target.value)} required/>
          <label>Enter your sex</label>
        </div>
        <div className={s.input_field}>
          <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required/>
          <label>Enter your email</label>
        </div>
        <div className={s.input_field}>
          <input type="tel" value={phone} onChange={(e) => setPhone(e.target.value)} required/>
          <label>Enter your phone</label>
        </div>
        <div className={s.input_field}>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required/>
          <label>Enter your password</label>
        </div>
        <div className={s.input_field}>
          <input type="password" value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)} required/>
          <label>Confirm your password</label>
        </div>
        <button type="button" id={s.btn} onClick={handleRegistration}>Create an account</button>
        <div className={s.register}>
          <p>Already have an account? <Link to={`/sign-in`}>Login</Link></p>
        </div>
      </form>
    </div>
    </div>
    // <div>
    //   <h2>Регистрация</h2>
    //   <form>
    //     <label>
    //       Имя:
    //       <input type="text" name="first_name" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
    //     </label>
    //     <br />
    //     <label>
    //       Фамилия:
    //       <input type="text" name="last_name" value={lastName} onChange={(e) => setLastName(e.target.value)} />
    //     </label>
    //     <br />
    //     <label>
    //       Пол:
    //       <input type="text" name="sex" value={sex} onChange={(e) => setSex(e.target.value)} />
    //     </label>
    //     <br />
    //     <label>
    //       Email:
    //       <input type="email" name="email" value={email} onChange={(e) => setEmail(e.target.value)} />
    //     </label>
    //     <br />
    //     <label>
    //       Телефон:
    //       <input type="tel" name="phone" value={phone} onChange={(e) => setPhone(e.target.value)} />
    //     </label>
    //     <br />
    //     <label>
    //       Пароль:
    //       <input
    //         type="password"
    //         name="password"
    //         value={password}
    //         onChange={(e) => setPassword(e.target.value)}
    //       />
    //     </label>
    //     <br />
    //     <label>
    //       Подтвердите пароль:
    //       <input
    //         type="password"
    //         name="confirmPassword"
    //         value={confirmPassword}
    //         onChange={(e) => setConfirmPassword(e.target.value)}
    //       />
    //     </label>
    //     <br />
    //     <button type="button" onClick={handleRegistration}>
    //       Зарегистрироваться
    //     </button>
    //   </form>
    // </div>
  );
};

export default RegistrationComponent;
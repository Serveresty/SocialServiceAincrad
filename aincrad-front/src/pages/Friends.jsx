import React, { useState, useEffect } from 'react';
import Cookies from 'universal-cookie';
import { useNavigate } from 'react-router-dom';
import s from '../styles/friends.module.css'
import loupe from '../static/loupe.svg'

const FriendsComponent = () => {
  const [responseData, setResponseData] = useState(null);
  const [friendCount, setFriendCount] = useState(0);
  const cookies = new Cookies();
  const navigate = useNavigate();

  useEffect(() => {
    const authToken = cookies.get('authToken');
    if (!authToken) {
      navigate('/sign-in');
    }
  }, [cookies, navigate]);

  useEffect(() => {
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');
    const section = urlParams.get('section');
    const backendUrl = 'http://localhost:8080/friends';

    const authToken = cookies.get('authToken');
    if (!authToken) {
      console.error('No authToken found in Cookie');
      return;
    }

    const requestOptions = {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `${authToken}`,
      },
    };

    const queryParams = id && section ? `?id=${id}&section=${section}` : '';

    const url = `${backendUrl}${queryParams}`;

    fetch(url, requestOptions)
      .then(response => response.json())
      .then(data => {
        setResponseData(data);
        if (data && data.data) {
          setFriendCount(data.data.length);
        }
      })
      .catch(error => console.error('Error:', error));
  }, []);

  return (
    <div>
      {responseData && (
        <div className={s.friendlist}>
          <div className={s.central}>
            <label>All friends: {friendCount}</label>
            <div className={s.search}>
              <input type='text' placeholder='Search friends'></input>
              <button id={s.sBtn}><img src={loupe} width="25" height="25" className={s.loupe}/></button>
            </div>
            <ul id={s.lst} style={{ listStyleType: 'none'}}>
              {responseData.data && responseData.data.map((friend, index) => (
                <li key={index}>
                  <button className={s.btn} onClick={() => handleButtonClick(friend)}>
                    <div className={s.avatar}>
                      <img src={`data:image/jpeg;base64,${friend.avatar}`} style={{ width: '150px', height: 'auto' }}/>
                    </div>
                    <div className={s.names}>{friend.friend_first_name} {friend.friend_last_name}</div>
                  </button>
                  <hr align="center" width="600" size="3" color="#0000dd" />
                </li>
              ))}
            </ul>
          </div>

          <div className={s.friendBtns}> 
            <button className={s.rightBtn}>My friends</button>
            <button className={s.rightBtn}>Friend requests</button>
            <button className={s.rightBtn}>Search friends</button>
          </div>

        </div>
      )}
    </div>
  );

  function handleButtonClick(friend) {
    // Здесь добавьте код для обработки нажатия кнопки
    console.log('Button clicked for:', friend);
  }
};

export default FriendsComponent;

import React, { useState, useEffect } from 'react';
import Cookies from 'universal-cookie';
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';
import s from '../styles/profile.module.css'

const ProfileComponent = () => {
    const [profileData, setProfileData] = useState(null);
    const [friendList, setFriendList] = useState(null);
    const { id } = useParams();
    //////////////////////////////
    const cookies = new Cookies();
    const history = useNavigate();
  
    useEffect(() => {
        const authToken = cookies.get('authToken');
        if (!authToken) {
        history('/sign-in');
        }
    }, [cookies, history]);
    //////////////////////////////

    useEffect(() => {
      const fetchProfileData = async () => {
        try {
          const backendUrl = 'http://localhost:8080'; // Изменено: URL должен быть базовым адресом
          const authToken = cookies.get('authToken');
    
          if (!authToken) {
            console.error('No authToken found in Cookie');
            return;
          }
    
          const response = await axios.get(`${backendUrl}/${id}`, {
            headers: {
              Authorization: authToken, // Изменено: передача токена без лишних символов
            },
          });
    
          // Проверка на успешный статус ответа
          if (response.status === 200) {
            setProfileData(response.data.data); // Изменено: взятие данных напрямую из ответа
          } else {
            console.error('Error fetching profile data:', response.statusText);
          }
        } catch (error) {
          console.error('Error fetching profile data:', error.message);
        }
      };
    
      fetchProfileData();
    }, [id]);
    

      const addFriend = () => {
        if (profileData) {
          const authToken = cookies.get('authToken');
          if (!authToken) {
            console.error('No authToken found in Cookie');
            return;
          }

          fetch(`http://localhost:8080/add-friend/${profileData.id}`, {
            method: 'POST',
            headers: {
              Authorization: `${authToken}`,
            },
          })
            .then(response => {
              if (!response.ok) {
                throw new Error('Network response was not ok');
              }
              return response.json();
            })
            .then(data => {
              console.log('Success:', data);
            })
            .catch(error => {
              console.error('Error:', error);
            });
        } else {
          console.log('Profile data is not loaded');
        }
      };

      const editProfile = () => {

      };
    
      return (
        <div>
          <div className={s.profileHead}>
            <div className={s.avatar}>
              {profileData && (
                <img src={`data:image/jpeg;base64,${profileData.avatar}`} style={{ width: '200px', height: 'auto' }}/>
              )}
            </div>
            <div className={s.general}>
              <div className={s.inf}>
                {profileData && (
                  <h1>{profileData.first_name} {profileData.last_name}</h1>
                )}
                {profileData && (
                  <label className={s.quote}>{profileData.quote}</label>
                )}
                <button type="button" className={s.more}>More info</button>
              </div>
              <div className={s.btns}>
              {profileData && profileData.mine === false && (
                <button type="submit" className={s.addButton} onClick={addFriend}>
                  Add friend
                </button>
              )}
              {profileData && profileData.mine === true && (
                <button type="submit" className={s.addButton} onClick={editProfile}>
                  Edit profile
                </button>
              )}
              </div>
            </div>
          </div>
            <div className={s.posts}>
              
            </div>
            <div className={s.friendList}>
              
            </div>
        </div>
      );
};

export default ProfileComponent;
import React, { useState, useEffect } from 'react';
import Cookies from 'universal-cookie';
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';
import s from '../styles/profile.module.css'
import like from '../static/like.svg'
import comment from '../static/comment.svg'
import views from '../static/views.svg'
import pin from '../static/pin.svg'

const ProfileComponent = () => {
    const [posts, setPosts] = useState(null)
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


            <div className={s.container}>
            <div className={s.ff}>

            <div className={s.newPost}>
                      <textarea className={s.postT} placeholder="Enter post data"></textarea>
                      <button type="submit" className={s.sendBtn}>Send</button>
                      <img src={pin} width="30" height="30" className={s.pin}/>
            </div>

            <div className={s.posts}>
                    <div className={s.post}>
                        <div className={s.datas}>
                            <div className={s.avatars}>
                              {profileData && (
                                <img src={`data:image/jpeg;base64,${profileData.avatar}`} style={{ width: '70px', height: 'auto' }}/>
                              )}
                            </div>
                            {profileData && (
                              <h1>{profileData.first_name} {profileData.last_name}</h1>
                            )}
                            </div>
                          <div className={s.postText}>
                            Народ, до 11 числа я в Москве!
                            Стримов не будет, но буду на съёмках, паре мероприятий, так что поделюсь интересностями.
                          </div>
                          <div className={s.pstIc}>
                            <img src={like} width="30" height="30"/> <text id={s.stats}>5</text>       
                            <img src={comment} width="23" height="23"/> <text id={s.stats}>7</text>
                            <img src={views} width="23" height="23"/> <text id={s.statsV}>14</text>
                          </div>
                    </div>
            </div>
            </div>

            <div className={s.right}>
            <div className={s.friendList}>
                <h1>Friends:</h1>
                <p>You don't have any friends yet</p>
            </div>
            <div className={s.presents}>
                <h1>Presents:</h1>
                <p>You don't have any gifts yet</p>
            </div>
            <div className={s.subs}>
                <h1>Subscriptions:</h1>
                <p>You haven't subscribed to anyone yet</p>
            </div>
            </div>

            </div>


        </div>
      );
};

export default ProfileComponent;
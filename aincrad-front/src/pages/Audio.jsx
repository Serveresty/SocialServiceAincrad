import React, { useState, useEffect } from 'react';
import Cookies from 'universal-cookie';
import { useParams, useNavigate } from 'react-router-dom';
import '../styles/audio_page.css'

const AudioGETComponent = () => {
    const [responseData, setResponseData] = useState(null);
    const cookies = new Cookies();
    const history = useNavigate();
  
    useEffect(() => {
        const authToken = cookies.get('authToken');
        if (!authToken) {
        history('/sign-in');
        }
    }, [cookies, history]);

    //const params = useParams();
    //const id = params.id || '';
    //const section = params.section || '';
    useEffect(() => {
        const urlParams = new URLSearchParams(window.location.search);
        const id = urlParams.get('id');
        // Определите URL вашего бэкенда
        const backendUrl = 'http://localhost:8080/audio';

        const authToken = cookies.get('authToken');
        if (!authToken) {
            console.error('No authToken found in Cookie');
            return;
        }
    
        // Опции для запроса
        const requestOptions = {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `${authToken}`,
          },
        };
    
        const queryParams = `?id=${id}`;

        const url = `${backendUrl}${queryParams}`;
    
        // Выполнение запроса
        fetch(url, requestOptions)
          .then(response => response.json())
          .then(data => {
            // Обработка данных от бэкенда
            setResponseData(data);
          })
          .catch(error => console.error('Error:', error));
      }, []);

      return (
          <div className="audio">
            {responseData && responseData.data && responseData.data.map(item => (
              <div key={item.id} className="item">
                <button className="button">
                  {item.name} - {item.author}
                </button>
                <div className="hidden-field">{item.id}</div>
              </div>
            ))}
          </div>
      );
};

export default AudioGETComponent;
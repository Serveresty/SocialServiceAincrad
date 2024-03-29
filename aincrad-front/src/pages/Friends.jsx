import React, { useState, useEffect } from 'react';
import Cookies from 'universal-cookie';
import { useParams, useNavigate } from 'react-router-dom';

const FriendsComponent = () => {
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
        const section = urlParams.get('section');
        // Определите URL вашего бэкенда
        const backendUrl = 'http://localhost:8080/friends';

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
    
        const queryParams = id && section ? `?id=${id}&section=${section}` : '';

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
        <div>
          {/* Вывод данных от бэкенда */}
          {responseData && (
            <div>
              <h2>Friendlist:</h2>
              <pre>{JSON.stringify(responseData, null, 2)}</pre>
            </div>
          )}
        </div>
      );
};

export default FriendsComponent;
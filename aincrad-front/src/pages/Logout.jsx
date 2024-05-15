import React, { useEffect } from 'react';
import Cookies from 'universal-cookie';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

const LogoutComponent = () => {
    const cookies = new Cookies();
    const history = useNavigate();

    useEffect(() => {
        const authToken = cookies.get('authToken');
        if (!authToken) {
          history('/sign-in');
        }
    }, [cookies, history]);

    useEffect(() => {
        const fetchLogoutData = async () => {
            try {
              const backendUrl = `http://localhost:8080/logout`;
      
              const authToken = cookies.get('authToken');
              if (!authToken) {
                console.error('No authToken found in Cookie');
                return;
              }
      
              const requestOptions = {
                headers: {
                  'Content-Type': 'application/json',
                  Authorization: `${authToken}`,
                },
              };
      
              const response = await axios.get(backendUrl, requestOptions);
              console.log(response.data.message);
            } catch (error) {
              console.error('Error fetching data:', error);
            }
          };
      
          fetchLogoutData();
          cookies.remove('authToken')
          history('/sign-in');
    }, [cookies, history]);

    return (
        <div></div>
      );
};

export default LogoutComponent;
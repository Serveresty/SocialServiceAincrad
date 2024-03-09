import React, { useState, useEffect } from 'react';
import Cookies from 'universal-cookie';
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';

const ProfileComponent = () => {
    const [profileData, setProfileData] = useState(null);
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
            const backendUrl = 'http://localhost:8080';
    
            const authToken = cookies.get('authToken');
    
            if (!authToken) {
              console.error('No authToken found in Cookie');
              return;
            }
    
            const response = await axios.get(`${backendUrl}/${id}`, {
              headers: {
                Authorization: `${authToken}`,
              },
            });
    
            setProfileData(response.data.data);
          } catch (error) {
            console.error('Error fetching profile data:', error);
          }
        };
    
        fetchProfileData();
      }, [id]);
    
      return (
        <div>
          <h1>Profile Data</h1>
          {profileData ? (
            <pre>{JSON.stringify(profileData, null, 2)}</pre>
          ) : (
            <p>Loading...</p>
          )}
        </div>
      );
};

export default ProfileComponent;
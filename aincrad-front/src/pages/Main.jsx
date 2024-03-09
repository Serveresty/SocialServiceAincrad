// EmptyComponent.js
import React, { useEffect } from 'react';
import Cookies from 'universal-cookie';
import { useNavigate } from 'react-router-dom';

const EmptyComponent = () => {
  const cookies = new Cookies();
    const history = useNavigate();
  
    useEffect(() => {
        const authToken = cookies.get('authToken');
        if (!authToken) {
        history('/sign-in');
        }
    }, [cookies, history]);

  return (
    <div>
    </div>
  );
};

export default EmptyComponent;
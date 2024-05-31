import React, { useState, useEffect } from 'react';
import Cookies from 'universal-cookie';
import { useParams, useNavigate } from 'react-router-dom';
//import './PopupWithInputs.css'; // Импортируем файл стилей

function UploadVideo() {
  const [isOpen, setIsOpen] = useState(false);
  const [inputValue1, setInputValue1] = useState('');
  const [inputValue2, setInputValue2] = useState('');
  const [file, setFile] = useState(null);

  const cookies = new Cookies();
    const history = useNavigate();

    useEffect(() => {
        const authToken = cookies.get('authToken');
        if (!authToken) {
        history('/sign-in');
        }
    }, [cookies, history]);

  const togglePopup = () => {
    setIsOpen(!isOpen);
  };

  const handleInputChange1 = (e) => {
    setInputValue1(e.target.value);
  };

  const handleInputChange2 = (e) => {
    setInputValue2(e.target.value);
  };

  const handleFileChange = (e) => {
    setFile(e.target.files[0]);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const authToken = cookies.get('authToken');
    if (!authToken) {
        console.error('No authToken found in Cookie');
        return;
    }

    const formData = new FormData();
    formData.append('title', inputValue1);
    formData.append('description', inputValue2);
    formData.append('video', file);

    try {
      const response = await fetch('http://localhost:8080/video/upload', {
        method: 'POST',
        headers: {
            Authorization: `${authToken}`,
          },
        body: formData,
      });

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }

      const result = await response.json();
      console.log('Success:', result);
    } catch (error) {
      console.error('Error:dasd asdasdasda', error);
    }

    togglePopup();
  };

  return (
    <div>
      <button onClick={togglePopup}>Открыть всплывающее окно</button>
      {isOpen && (
        <div className="popup-container">
          <div className="popup-content">
            <button className="close-btn" onClick={togglePopup}>Закрыть</button>
            <form onSubmit={handleSubmit}>
              <input
                type="text"
                value={inputValue1}
                onChange={handleInputChange1}
                placeholder="Введите название"
              />
              <input
                type="text"
                value={inputValue2}
                onChange={handleInputChange2}
                placeholder="Введите описание"
              />
              <input
                type="file"
                onChange={handleFileChange}
              />
              <button type="submit">Отправить</button>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}

export default UploadVideo;
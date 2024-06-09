import React, { useState, useEffect } from 'react';
import Cookies from 'universal-cookie';
import { useParams, useNavigate } from 'react-router-dom';
import s from '../styles/audio_upload.module.css'
import u from '../static/upload.svg'
//import './PopupWithInputs.css'; // Импортируем файл стилей

function PopupWithInputs() {
  const [isOpen, setIsOpen] = useState(false);
  const [inputValue1, setInputValue1] = useState('');
  const [inputValue2, setInputValue2] = useState('');
  const [file, setFile] = useState(null);
  const [fileName, setFileName] = useState('');

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
    const file = e.target.files[0];
    setFile(file);
    setFileName(file ? file.name : '');
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
    formData.append('author', inputValue2);
    formData.append('audio', file);

    try {
      const response = await fetch('http://localhost:8080/audio/upload', {
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
    <div className={s.btttn}>
      <button className={s.opn} onClick={togglePopup}><img src={u} width="20" height="20"/></button>
      {isOpen && (
        <div className={s.popupContainer}>
          <div className={s.popupContent}>
            <button className={s.closeButton} onClick={togglePopup}>×</button>
            <form onSubmit={handleSubmit}>
              <label className={s.lbl}>Введите название: </label>
              <input className={s.inpt_name}
                type="text"
                value={inputValue1}
                onChange={handleInputChange1}
                placeholder="Введите название"
              />
              <label className={s.lbl}>Введите исполнителя: </label>
              <input className={s.inpt_desc}
                type="text"
                value={inputValue2}
                onChange={handleInputChange2}
                placeholder="Введите исполнителя"
              />
              <label className={s.custom_file_upload}>
                Загрузить файл
                <input id={s.file_upload} type="file" onChange={handleFileChange} />
              </label>
              {fileName && <span className={s.file_name}>{fileName}</span>}
              <button type="submit" className={s.send_vid}>Отправить</button>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}

export default PopupWithInputs;

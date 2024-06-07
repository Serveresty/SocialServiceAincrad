import React, { useState, useEffect } from 'react';
import Cookies from 'universal-cookie';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import PopupWithInputs from './UploadAudio';
import s from '../styles/audio.module.css'

const AudioGETComponent = () => {
  const cookies = new Cookies();
  const history = useNavigate();
  const [responseData, setResponseData] = useState([]);
  const [audioElement, setAudioElement] = useState(null);
  const [currentAudioUrl, setCurrentAudioUrl] = useState(null);
  const [isPlaying, setIsPlaying] = useState(false);
  const [currentId, setCurrentId] = useState(0);

  useEffect(() => {
    const authToken = cookies.get('authToken');
    if (!authToken) {
      history('/sign-in');
    }
  }, [cookies, history]);

  useEffect(() => {
    const fetchAudioData = async () => {
      try {
        const urlParams = new URLSearchParams(window.location.search);
        const id = urlParams.get('id');
        const backendUrl = `http://localhost:8080/audio?id=${id}`;

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
        setResponseData(response.data.data);
      } catch (error) {
        console.error('Error fetching audio data:', error);
      }
    };

    fetchAudioData();
  }, [cookies, currentId]);

  const handlePlayAudio = async (id) => {
    try {
      const authToken = cookies.get('authToken');
      if (!authToken) {
        console.error('No authToken found in Cookie');
        return;
      }

      const backendUrl = `http://localhost:8080/audio/${id}`;
      const requestOptions = {
        headers: {
          Authorization: `${authToken}`,
        },
        responseType: 'blob',
      };

      const response = await axios.get(backendUrl, requestOptions);
      const audioBlob = new Blob([response.data], { type: 'audio/mp3' });
      const audioUrl = URL.createObjectURL(audioBlob);
      console.log(response.data)

      if (currentId !== id || currentId == 0) {
        if (audioElement !== null) {
          const currentTime = audioElement.currentTime;
          audioElement.pause();
          audioElement.currentTime = 0;
          audioElement.src = audioUrl;
          audioElement.play();
          setCurrentAudioUrl(audioUrl);
          setCurrentId(id);
          setIsPlaying(true);
        } else {
          const newAudioElement = new Audio(audioUrl);
          newAudioElement.play();
          setAudioElement(newAudioElement);
          setCurrentAudioUrl(audioUrl);
          setCurrentId(id);
          setIsPlaying(true);
        }
      } else {
        if (isPlaying) {
          audioElement.pause();
        } else {
          audioElement.play();
        }
        setIsPlaying(!isPlaying);
      }
    } catch (error) {
      console.error('Error playing audio:', error);
    }
  };

  return (
    <div>
      <div><PopupWithInputs/></div>
      <div className={s.audio_container}>
        <ul id={s.audios}>
        {Array.isArray(responseData) && responseData.map(audio => (
          <li key={audio.id}><button onClick={() => handlePlayAudio(audio.id)}>
            {`${audio.name} - ${audio.author}`}
          </button>
          </li>
        ))}
        </ul>
      </div>
    </div>
  );
};

export default AudioGETComponent;

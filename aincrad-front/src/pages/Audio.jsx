import React, { useState, useEffect, useRef } from 'react';
import Cookies from 'universal-cookie';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

const AudioGETComponent = () => {
  const [responseData, setResponseData] = useState(null);
  const [isPlaying, setIsPlaying] = useState([]);
  const audioRef = useRef(new Audio());
  const [currentAudio, setCurrentAudio] = useState(null);
  const [currentAudio1, setCurrentAudio1] = useState(null);

  const cookies = new Cookies();
  const history = useNavigate();

  useEffect(() => {
    const authToken = cookies.get('authToken');
    if (!authToken) {
      history('/sign-in');
    }
  }, [cookies, history]);

  const handleUserGesture = () => {
    const audioContext = new (window.AudioContext || window.webkitAudioContext)();
    if (audioContext.state === 'suspended') {
      audioContext.resume();
    }
  };

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

        const config = {
          headers: {
            'Content-Type': 'application/json',
            Authorization: authToken,
          },
        };

        const response = await axios.get(backendUrl, config);
        setResponseData(response.data);

        if (response.data.data.length > 0) {
          preloadAudio(response.data.data);
          setIsPlaying(Array(response.data.data.length).fill(false));
        }
      } catch (error) {
        console.error('Error fetching audio data:', error);
      }
    };

    fetchAudioData();
  }, [cookies]);

  const preloadAudio = (audioList) => {
    audioList.forEach(audio => {
      const audioElement = audioRef.current;

      const source = document.createElement('source');
      source.src = `http://localhost:8080/audio/${audio.id}.mp3`;
      source.type = 'audio/mp3';

      audioElement.appendChild(source);
    });
  };

  const handleAudioButtonClick = async (id) => {
    const audioElement = audioRef.current;
    const index = responseData.data.findIndex(audio => audio.id === id);

    if (audioElement) {
      if (isPlaying[index]) {
        audioElement.pause();
      } else {
        if (currentAudio === `http://localhost:8080/audio/${id}`) {
          // Если выбрана та же песня, возобновляем воспроизведение
          audioElement.play();
        } else {
          // В противном случае, устанавливаем новую песню
          audioElement.src = `http://localhost:8080/audio/${id}`;
          audioElement.play();
          setCurrentAudio1(responseData.data.find(audio => audio.id === id));
        }
      }

      setCurrentAudio(`http://localhost:8080/audio/${id}`);
      setIsPlaying(prevState => {
        const newState = [...prevState];
        newState.fill(false);
        newState[index] = !prevState[index];
        return newState;
      });
    }
  };

  const handlePauseButtonClick = async () => {
    const audioElement = audioRef.current;

    if (audioElement && !audioElement.paused) {
      audioElement.pause();
      setIsPlaying(Array(responseData.data.length).fill(false));
    }
  };

  return (
    <div>
      {audioRef.current && (
        <div>
          Now Playing:
          {currentAudio && (
            <label> {currentAudio1.author} - {currentAudio1.name}</label>
          )}
          <p/>
          <button onClick={handlePauseButtonClick}>
            Pause
          </button>
        </div>
      )}
      
      {responseData && responseData.data && responseData.data.map(audio => (
        <div key={audio.id}>
          <p>{audio.name} - {audio.author}</p>
          <button onClick={() => handleAudioButtonClick(audio.id)}>
            {isPlaying[audio.id] ? 'Pause' : 'Play'}
          </button>
        </div>
      ))}
    </div>
  );
};

export default AudioGETComponent;

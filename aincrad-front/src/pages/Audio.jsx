import React, { useState, useEffect, useRef } from 'react';
import Cookies from 'universal-cookie';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import PopupWithInputs from './UploadAudio';
import s from '../styles/audio.module.css';
import note from '../static/note.svg'

const AudioGETComponent = () => {
  const cookies = new Cookies();
  const navigate = useNavigate();
  const [responseData, setResponseData] = useState([]);
  const [currentAudioUrl, setCurrentAudioUrl] = useState(null);
  const [isPlaying, setIsPlaying] = useState(false);
  const [currentId, setCurrentId] = useState(0);
  const [currentName, setCurrentName] = useState('');
  const [currentTime, setCurrentTime] = useState(0);
  const [duration, setDuration] = useState(0);
  const audioRef = useRef(null);

  useEffect(() => {
    const authToken = cookies.get('authToken');
    if (!authToken) {
      navigate('/sign-in');
    }
  }, [cookies, navigate]);

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
  }, [cookies]);

  const handlePlayAudio = async (id, name) => {
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

      if (currentId !== id) {
        if (audioRef.current) {
          audioRef.current.pause();
          audioRef.current.src = audioUrl;
          audioRef.current.play();
          setCurrentAudioUrl(audioUrl);
          setCurrentId(id);
          setCurrentName(name);
          setIsPlaying(true);
        }
      } else {
        if (isPlaying) {
          audioRef.current.pause();
        } else {
          audioRef.current.play();
        }
        setIsPlaying(!isPlaying);
      }
    } catch (error) {
      console.error('Error playing audio:', error);
    }
  };

  const formatTime = (time) => {
    const minutes = Math.floor(time / 60);
    const seconds = Math.floor(time % 60);
    return `${minutes}:${seconds < 10 ? '0' : ''}${seconds}`;
  };

  useEffect(() => {
    const audio = audioRef.current;
    if (audio) {
      const updateTime = () => setCurrentTime(audio.currentTime);
      const updateDuration = () => setDuration(audio.duration);

      audio.addEventListener('timeupdate', updateTime);
      audio.addEventListener('loadedmetadata', updateDuration);

      return () => {
        audio.removeEventListener('timeupdate', updateTime);
        audio.removeEventListener('loadedmetadata', updateDuration);
      };
    }
  }, [audioRef.current]);

  return (
    <div className={s.msCont}>
      <div>
      <div className={s.controll}>
        <div className={s.audio_player}>
          <div className={s.timeline}>
            <div
              className={s.progress}
              style={{ width: `${(currentTime / duration) * 100}%` }}
            ></div>
          </div>
          <div className={s.controls}>
            <div
              className={s.play_container}
              onClick={() => isPlaying ? audioRef.current.pause() : audioRef.current.play()}
            >
              <div className={`${s.toggle_play} ${isPlaying ? s.pause : s.play}`}></div>
            </div>
            <div className={s.time}>
              <div className={s.current}>{formatTime(currentTime)}</div>
              <div className={s.divider}>/</div>
              <div className={s.length}>{formatTime(duration)}</div>
            </div>
            <div className={s.name}>{currentName}</div>
            <div className={s.volume_container}>
              <div className={s.volume_button}>
                <div className={s.volume}></div>
              </div>
              <div className={s.volume_slider}>
                <div className={s.volume_percentage}></div>
              </div>
            </div>
          </div>
        </div>
        <PopupWithInputs />
      </div>
      <div className={s.audio_container}>
        <ul id={s.audios}>
          {Array.isArray(responseData) && responseData.map(audio => (
            <li key={audio.id}>
              <button className={s.ms} onClick={() => handlePlayAudio(audio.id, audio.name)}>
              <img src={note} width="20" height="20"/> {`${audio.name} - ${audio.author}`}
              </button>
            </li>
          ))}
        </ul>
      </div>
      <audio ref={audioRef} onPlay={() => setIsPlaying(true)} onPause={() => setIsPlaying(false)} />
    </div>
          <div className={s.friendList}>
              <h1>Friends:</h1>
              <p>You don't have any friends yet</p>
          </div>
    </div>
  );
};

export default AudioGETComponent;

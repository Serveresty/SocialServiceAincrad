import React from 'react';
import { createBrowserRouter, RouterProvider, BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import MainComponent from './pages/Main';
import ProfileComponent from './pages/Profile';
import FriendsComponent from './pages/Friends';
import AudioGETComponent from './pages/Audio';
import Header from './components/Header';
import LeftBar from './components/LeftBar';
import Chat from './pages/Chat';
import LogoutComponent from './pages/Logout';
import VideoGrid from './pages/Video';
import s from './styles/main.module.css'

function MainLayout() {
    return (
        <div className={s.backgroundd}>
            <Header />
            <div className={s.sec_container}>
            <LeftBar />
            <div className={s.main_container}>
            <Routes>
                <Route path="/" element={<MainComponent />}/>
                <Route path="/logout" element={<LogoutComponent />} />
                <Route path="/:id" element={<ProfileComponent />} />
                <Route path="/friends" element={<FriendsComponent />} />
                <Route path="/audio" element={<AudioGETComponent />} />
                <Route path="/messages" element={<Chat />} />
                <Route path="/video/:id" element={<VideoGrid />} />
            </Routes>
            </div>
            </div>
        </div>
    );
  }
  
export default MainLayout;
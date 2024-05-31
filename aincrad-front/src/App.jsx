import { createBrowserRouter, RouterProvider, BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import { useState } from 'react'
import { Filecontext } from './contexts/Filecontext';
import './App.css'
import AuthComponent from './pages/Auth';
import RegistrationComponent from './pages/Registration';
import EmptyComponent from './pages/Main';
import ProfileComponent from './pages/Profile';
import FriendsComponent from './pages/Friends';
import AudioGETComponent from './pages/Audio';
import Header from './components/Header';
import LeftBar from './components/LeftBar';
import Chat from './pages/Chat';
import LogoutComponent from './pages/Logout';
import VideoGrid from './pages/Video';

function App() {
  const [logID, setLogID] = useState("")
  return (
    <div>
      <Filecontext.Provider value={{ logID, setLogID }}>
        <Header />
        <div className="main-container">
        <LeftBar />
        <Routes>
          <Route path="/" element={<EmptyComponent />}/>
          <Route path="/sign-up" element={<RegistrationComponent />} />
          <Route path="/sign-in" element={<AuthComponent />} />
          <Route path="/logout" element={<LogoutComponent />} />
          <Route path="/:id" element={<ProfileComponent />} />
          <Route path="/friends" element={<FriendsComponent />} />
          <Route path="/audio" element={<AudioGETComponent />} />
          <Route path="/messages" element={<Chat />} />
          <Route path="/video/:id" element={<VideoGrid />} />
        </Routes>
        </div>
      </Filecontext.Provider>
    </div>
  );
};

export default App

import { createBrowserRouter, RouterProvider, BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import { useState } from 'react'
import { Filecontext } from './contexts/Filecontext';
import './App.css'
import AuthComponent from './pages/Auth';
import RegistrationComponent from './pages/Registration';
import MainLayout from './MainApp';

function App() {
  const [logID, setLogID] = useState("")
  return (
    <div>
      <Filecontext.Provider value={{ logID, setLogID }}>
        <Routes>
          <Route path="*" element={<MainLayout />}/>
          <Route path="/sign-up" element={<RegistrationComponent />} />
          <Route path="/sign-in" element={<AuthComponent />} />
        </Routes>
      </Filecontext.Provider>
    </div>
  );
};

export default App

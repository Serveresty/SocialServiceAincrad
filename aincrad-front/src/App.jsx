import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import { createBrowserRouter, RouterProvider, Route, Link } from 'react-router-dom';
import './App.css'
import AuthComponent from './pages/Auth';
import RegistrationComponent from './pages/Registration';

const router = createBrowserRouter([
  {
    path: "/sign-up",
    element: <RegistrationComponent/>
  },
  {
    path: "/sign-in",
    element: <AuthComponent/>,
  },
]);

function App() {
  return (
    <div>
        <RouterProvider router={router} />
    </div>
  );
};

export default App

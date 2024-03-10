import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import { createBrowserRouter, RouterProvider, Route, Link } from 'react-router-dom';
import './App.css'
import AuthComponent from './pages/Auth';
import RegistrationComponent from './pages/Registration';
import EmptyComponent from './pages/Main';
import ProfileComponent from './pages/Profile';
import FriendsComponent from './pages/Friends';

const router = createBrowserRouter([
  {
    path: "/",
    element: <EmptyComponent/>
  },
  {
    path: "/sign-up",
    element: <RegistrationComponent/>
  },
  {
    path: "/sign-in",
    element: <AuthComponent/>,
  },
  {
    path: "/:id",
    element: <ProfileComponent/>
  },
  {
    path: "/friends",
    element: <FriendsComponent/>
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

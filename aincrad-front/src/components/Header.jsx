import React from "react";
import { useContext } from 'react'
import { Filecontext } from '../contexts/Filecontext';
import { Link } from "react-router-dom";
import logo from '../logo_test.svg';
import '../styles/header.css';

const Header = () => {
    const { logID } = useContext(Filecontext)
    return (
        <header>
            <div className="header-container">
                <ul id="header-nav">
                    <li><button className="btns">Music</button></li>
                    <li><Link to="/"><img src={logo} height="50" width="200"></img></Link></li>
                    <li><Link to={`/${logID}`}><button className="btns">Profile</button></Link></li>
                </ul>
            </div>
        </header>
    )
};

export default Header;
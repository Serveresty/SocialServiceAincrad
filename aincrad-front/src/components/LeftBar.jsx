import React from "react";
import { Filecontext } from "../contexts/Filecontext";
import { useContext } from 'react'
import '../styles/left_bar.css';
import { Link } from "react-router-dom";

const LeftBar = () => {
    const { logID } = useContext(Filecontext)
    return (
        <div className="left-bar-container">
            <ul id="left-bar-nav">
                <li><Link to={`/${logID}`}>Profile</Link></li>
                <li><a href="/news">News</a></li>
                <li><a href="/message">Messages</a></li>
                <li><a href="/friends">Friends</a></li>
                <li><a href="/groups">Groups</a></li>
                <li><a href="/photos">Photos</a></li>
                <li><a href="/audio">Audios</a></li>
                <li><a href="/videos">Videos</a></li>
            </ul>
        </div>
    )
};

export default LeftBar;
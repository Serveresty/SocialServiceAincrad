import React from "react";
import '../styles/left_bar.css';

const LeftBar = () => {
    return (
        <div class="left-bar-container">
            <ul id="left-bar-nav">
                <li><a href="/profile">Profile</a></li>
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
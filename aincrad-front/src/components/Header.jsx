import React from "react";
import logo from '../logo_test.svg';
import '../styles/header.css';

const Header = () => {
    return (
        <header>
            <div class="header-container">
                <ul id="header-nav">
                    <li><button>Music</button></li>
                    <li><a href="/"><img src={logo} height="50" width="200"></img></a></li>
                    <li><button>Profile</button></li>
                </ul>
            </div>
        </header>
    )
};

export default Header;
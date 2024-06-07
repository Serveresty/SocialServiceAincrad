import React from 'react';
import s from '../styles/video_grid.module.css';

const VideoPreview = ({ title, created, preview, views }) => {
    return (
        <div className={s.video_preview}>
            <img src={`data:image/jpeg;base64,${preview}`} alt={`${title} preview`} />
            <h3>{title}</h3>
            <p>Views: {views}</p>
            <p>Created at: {created}</p>
        </div>
    );
};

export default VideoPreview;
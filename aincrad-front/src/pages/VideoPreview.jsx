import React from 'react';

const VideoPreview = ({ title, created, preview, views }) => {
    return (
        <div className="video-preview">
            <img src={`data:image/jpeg;base64,${preview}`} alt={`${title} preview`} />
            <h3>{title}</h3>
            <p>Views: {views}</p>
            <p>Created at: {created}</p>
        </div>
    );
};

export default VideoPreview;
import React, { useEffect, useState } from 'react';
import Cookies from 'universal-cookie';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';
import VideoPreview from './VideoPreview'
import UploadVideo from './UploadVideo';
// import './VideoGrid.css';

const VideoGrid = () => {
    const cookies = new Cookies();
    const history = useNavigate();
    const [videos, setVideos] = useState([]);
    const { id } = useParams();

    useEffect(() => {
        const authToken = cookies.get('authToken');
        if (!authToken) {
          history('/sign-in');
        }
      }, [cookies, history]);

    useEffect(() => {
        const fetchVideos = async () => {
            try {
                const backendUrl = 'http://localhost:8080/video';

                const authToken = cookies.get('authToken');
    
                if (!authToken) {
                console.error('No authToken found in Cookie');
                return;
                }
                
                const response = await axios.get(`${backendUrl}/${id}`, {
                    headers: {
                      Authorization: `${authToken}`,
                    },
                  });
                console.log(response.data.data)
                setVideos(response.data.data);
            } catch (error) {
                console.error("Error fetching videos", error);
            }
        };

        fetchVideos();
    }, [id]);

    return (
        <div>
            <div><UploadVideo /></div>
            <div className="video-grid">
                {Array.isArray(videos) && videos.map(video => (
                    <VideoPreview 
                        key={video.id} 
                        title={video.title} 
                        created={video.created_at} 
                        preview={video.preview} 
                    />
                ))}
            </div>
        </div>
    );
};

export default VideoGrid;
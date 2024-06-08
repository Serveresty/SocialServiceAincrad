import React, { useEffect, useState } from 'react';
import Cookies from 'universal-cookie';
import { useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';
import VideoPreview from './VideoPreview'
import UploadVideo from './UploadVideo';
import s from '../styles/video_grid.module.css';

const VideoGrid = () => {
    const cookies = new Cookies();
    const history = useNavigate();
    const [videos, setVideos] = useState([]);
    const [modalOpen, setModalOpen] = useState(false);
    const [selectedVideo, setSelectedVideo] = useState(null);
    const [comment, setComment] = useState("");
    const [comments, setComments] = useState([]);
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

    const fetchVideoFile = async (userID, videoID) => {
        try {
            const authToken = cookies.get('authToken');

            if (!authToken) {
                console.error('No authToken found in Cookie');
                return;
            }

            const backendUrl = `http://localhost:8080/video/${userID}/${videoID}`;
            const requestOptions = {
                headers: {
                  Authorization: `${authToken}`,
                },
                responseType: 'blob',
              };
            const response = await axios.get(backendUrl, requestOptions);
            const videoBlob = new Blob([response.data], { type: 'video/mp4' });
            const videoUrl = URL.createObjectURL(videoBlob);
            return videoUrl;
        } catch (error) {
            console.error('Error fetching video file', error);
            return null;
        }
    };

    const handleVideoClick = async (video) => {
        try {
            const videoUrl = await fetchVideoFile(id, video.id);
            setSelectedVideo(videoUrl);
            // const commentsResponse = await axios.get(`http://localhost:8080/comments/${video.id}`, {
            //     headers: {
            //         Authorization: `${cookies.get('authToken')}`,
            //     },
            // });
            // setComments(commentsResponse.data.comments);
            setModalOpen(true);
        } catch (error) {
            console.error('Error handling video click', error);
        }
    };

    const closeModal = () => {
        setModalOpen(false);
    };

    const handleCommentSubmit = async (event) => {
        event.preventDefault();
        try {
            const authToken = cookies.get('authToken');
            const response = await axios.post(`http://localhost:8080/comments/${selectedVideo.id}`, 
            {
                comment: comment,
            }, 
            {
                headers: {
                    Authorization: `${authToken}`,
                },
            });
            setComments([...comments, response.data.comment]);
            setComment("");
        } catch (error) {
            console.error('Error submitting comment', error);
        }
    };

    return (
        <div>
            <div className={s.upload}><UploadVideo /></div>
            <div className={s.video_grid}>
                {Array.isArray(videos) && videos.map(video => (
                    <button className={s.vid}
                        key={video.id}
                        onClick={() => handleVideoClick(video)}
                    >
                        <VideoPreview 
                            key={video.id} 
                            title={video.title} 
                            created={video.created_at} 
                            preview={video.preview}
                            views={video.views} 
                        />
                    </button>
                ))}
            </div>
            {modalOpen && (
                <div className={s.modal_overlay} onClick={closeModal}>
                    <div className={s.modal} onClick={e => e.stopPropagation()}>
                        <button className={s.close_button} onClick={closeModal}>
                            &times;
                        </button>
                        <video controls autoPlay>
                            <source src={selectedVideo} type="video/mp4" />
                            Your browser does not support the video tag.
                        </video>
                        <div className={s.comments_section}>
                            <h3>Comments</h3>
                            <ul>
                                {comments.map((comment, index) => (
                                    <li key={index}>{comment.text}</li>
                                ))}
                            </ul>
                            <form onSubmit={handleCommentSubmit}>
                                <textarea
                                    value={comment}
                                    onChange={(e) => setComment(e.target.value)}
                                    placeholder="Leave a comment"
                                />
                                <button type="submit">Submit</button>
                            </form>
                        </div>
                    </div>
                </div>
            )}
        </div>
    );
};

export default VideoGrid;
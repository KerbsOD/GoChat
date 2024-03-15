import React from 'react'
import { useNavigate } from 'react-router-dom'
import './Login.css'

function Login({username, setUsername}) {
    const navigate = useNavigate()

    const onButtonClick = () => {
        if (username === '') {
            return
        }
        navigate('/chat')
    }

    const handleKeyDown = (event) => {
        if (event.key === 'Enter') {
            onButtonClick();
        }
    }

    return (
        <div className='background-image'>
            <div className={'mainContainer'}>
                <div className={'titleContainer'}>
                    <div>Welcome to GoChat!</div>
                </div>
                <br />

                <div className={'inputContainer'}>
                    <input
                        placeholder="Enter username"
                        onChange={(ev) => setUsername(ev.target.value)}
                        onKeyDown={handleKeyDown}
                        className={'inputBox'}
                    />
                </div>
                <br />

                <div className={'inputContainer'}>
                    <input className={'inputButton'} type="button" onClick={onButtonClick} value={'Enter'} />
                </div>
            </div>
        </div>
    )
}

export default Login

import React from 'react'
import { useNavigate } from 'react-router-dom'

export default function Home({username, setUsername}) {
    const navigate = useNavigate()

    const onButtonClick = () => {
        if (username === '') {
            return
        }
        navigate('/chat')
    }

    return (
        <div className={'mainContainer'}>
            <div className={'titleContainer'}>
                <div>Welcome to GoChat!</div>
            </div>
            <br />

            <div className={'inputContainer'}>
                <input
                    placeholder="Enter username"
                    onChange={(ev) => setUsername(ev.target.value)}
                    className={'inputBox'}
                />
            </div>
            <br />

            <div className={'inputContainer'}>
                <input className={'inputButton'} type="button" onClick={onButtonClick} value={'Enter'} />
            </div>
        </div>
    )
}

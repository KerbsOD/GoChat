import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Login from './Login.js'
import Chat from './Chat.js'
import './App.css'
import { useState } from 'react'

function App() {
    const [username, setUsername] = useState('')
    
    return (
        <div className="App">
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Login username={username} setUsername={setUsername} />} />
                    <Route path="/chat" element={<Chat username={username} />} />
                </Routes>
            </BrowserRouter>
        </div>
    )
}

export default App


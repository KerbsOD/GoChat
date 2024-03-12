import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Home from './Home.js'
import Chat from './Chat.js'
import './App.css'
import { useState } from 'react'

function App() {
    const [username, setUsername] = useState('')
    
    return (
        <div className="App">
            <BrowserRouter>
                <Routes>
                    <Route path="/" element={<Home username={username} setUsername={setUsername} />} />
                    <Route path="/chat" element={<Chat username={username} />} />
                </Routes>
            </BrowserRouter>
        </div>
    )
}

export default App

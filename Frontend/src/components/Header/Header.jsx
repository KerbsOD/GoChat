import "./Header.scss";
import { Link } from 'react-router-dom'; // Import Link from React Router

const Header = () => {
    return (
        <header className="Header">
            <Link to="/" className="backButton">&larr;</Link> {/* Move the Link component inside the div */}
            <h1>GoChat</h1>
        </header> 
    );
};

export default Header;
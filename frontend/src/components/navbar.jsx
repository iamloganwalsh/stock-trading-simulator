import React from 'react';
import { Link, useLocation } from 'react-router-dom';

const NavBar = () => {
  const location = useLocation();

  const styles = {
    nav: {
      position: "fixed",
      top: 0,
      left: 0,
      height: "100vh",
      width: "100px",
      backgroundColor: "#333",
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
      padding: "10px 0",
      boxShadow: "2px 0 5px rgba(0,0,0,0.2)",
    },
    link: {
      textDecoration: "none",
      color: "white",
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
      margin: "20px 0",
      fontSize: "12px",
      transition: "all 0.3s ease",
    },
    activeLink: {
      color: "#a3cfc5", 
      fontWeight: "bold",
    },
    icon: {
      width: "24px",
      height: "24px",
      marginBottom: "5px",
    },
  };

  const getLinkStyle = (path) => {
    return location.pathname === path
      ? { ...styles.link, ...styles.activeLink }
      : styles.link;
  };

  return (
    <nav style={styles.nav}>
      <Link to="/" style={getLinkStyle('/')}>
        <svg viewBox="0 0 24 24" style={styles.icon} fill="currentColor">
          <path d="M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z" />
        </svg>
        Home
      </Link>
      <Link to="/portfolio" style={getLinkStyle('/portfolio')}>
        <svg viewBox="0 0 24 24" style={styles.icon} fill="currentColor">
          <path d="M16 6l2.29 2.29-4.88 4.88-4-4L2 16.59 3.41 18l6-6 4 4 6.3-6.29L22 12V6z" />
        </svg>
        Portfolio
      </Link>
      <Link to="/trade" style={getLinkStyle('/trade')}>
        <svg viewBox="0 0 24 24" style={styles.icon} fill="currentColor">
          <path d="M8 9l-4 4 4 4v-3h8v3l4-4-4-4v3H8V9z" />
        </svg>
        Trade
      </Link>
      <Link to="/market" style={getLinkStyle('/market')}>
        <svg viewBox="0 0 24 24" style={styles.icon} fill="currentColor">
          <path d="M5 9.2h3V19H5zM10.6 5h2.8v14h-2.8zm5.6 8H19v6h-2.8z" />
        </svg>
        Market
      </Link>
      <Link to="/stocksportfolio" style={getLinkStyle('/stocksportfolio')}>
        <svg viewBox="0 0 24 24" style={styles.icon} fill="currentColor">
          <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-6h2v6zm0-8h-2V7h2v2z" />
        </svg>
        Stocks
      </Link>
      <Link to="/account" style={getLinkStyle('/account')}>
        <svg viewBox="0 0 24 24" style={styles.icon} fill="currentColor">
          <path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z" />
        </svg>
        Account
      </Link>
    </nav>
  );
};

export default NavBar;
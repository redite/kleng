import React from 'react';

const Navbar = ({ onToggleSidebar }) => {
  return (
    <nav className="navbar navbar-default">
      <div className="container-fluid">
        <div className="navbar-header">
          <a
            id="sidebarCollapse"
            className="btn btn-info navbar-btn"
            onClick={onToggleSidebar()}>
            <i className="glyphicon glyphicon-align-left" />
            Toggle Sidebar
          </a>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;

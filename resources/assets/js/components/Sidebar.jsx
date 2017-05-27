import React, { PropTypes } from 'react';

const Sidebar = ({ showSidebar }) => {
  let className = '';

  if (showSidebar) {
    className = 'active';
  }

  return (
    <nav id="sidebar" className={`${className}`}>
      <div className="sidebar-header">
        <h3>Kleng</h3>
      </div>
      <ul className="list-unstyled components">
        <p>Dummy Heading</p>
        <li>
          <a href="#">Home</a>
        </li>
        <li>
          <a href="#">About</a>
        </li>
      </ul>
    </nav>
  );
};

Sidebar.propTypes = {
  showSidebar: PropTypes.bool.isRequired
};

export default Sidebar;

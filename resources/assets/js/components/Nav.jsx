import React from 'react';

const Nav = () => {
  return (
    <div className="uk-navbar-container" data-uk-sticky="media: 960">
      <div className="uk-container uk-container-expand">
        <nav className="uk-navbar">
          <div className="uk-navbar-left">
            <a href="#" className="uk-navbar-item uk-logo">Kleng</a>
          </div>
          <div className="uk-navbar-right">
            <ul className="uk-navbar-nav uk-visible@m">
              <li><a href="#">About</a></li>
            </ul>
          </div>
        </nav>
      </div>
    </div>
  );
};

export default Nav;

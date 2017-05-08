import React, { Component } from 'react';
import { connect } from 'react-redux';
import Sidebar from '../components/Sidebar';
import Navbar from '../components/Navbar';
import { toggleSidebar } from '../actions';

class App extends Component {
  constructor(props) {
    super(props);

    this.onToggleSidebar = this.onToggleSidebar.bind(this);
  }

  render() {
    const { showSidebar } = this.props;

    return (
      <div>
        <title>Kleng</title>
        <div className="wrapper">
          <Sidebar showSidebar={showSidebar} />
          <div id="content">
            <Navbar onToggleSidebar={this.onToggleSidebar} />
          </div>
        </div>
      </div>
    );
  }

  onToggleSidebar() {
    return e => {
      e.preventDefault();

      this.props.onToggleSidebar();
    };
  }
}

const mapStateToProps = state => {
  const { showSidebar } = state.sidemenu;

  return {
    showSidebar
  };
}

const mapDispatchToProps = dispatch => {
  return {
    onToggleSidebar: () => {
      dispatch(toggleSidebar());
    }
  };
}

export default connect(mapStateToProps, mapDispatchToProps)(App);

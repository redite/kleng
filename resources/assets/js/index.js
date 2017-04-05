import React from 'react';
import ReactDOM from 'react-dom';
import App from './components/App';
import UIkit from 'uikit';
import Icons from 'uikit/dist/js/uikit-icons';

UIkit.use(Icons);

window.$ = window.jQuery = require('jquery');

ReactDOM.render(<App />, document.getElementById('app'));

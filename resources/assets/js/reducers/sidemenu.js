const initialState = {
  showSidebar: true
};

const sidemenu = (state = initialState, action) => {

  if (action.type === 'TOGGLE_SIDEBAR') {
    return {
      ...state,
      showSidebar: !state.showSidebar
    };
  }

  return state;
};

export default sidemenu;

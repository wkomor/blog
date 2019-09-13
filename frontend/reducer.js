
initialState = {
  posts: [],
  page: 1
}

export function reducer(state, action) {
    switch (action.type) {
  
      case FETCH_POSTS: {
        const {posts} = action.payload
        return {...state, posts}
      }
  
      default:
        return state
    }
  }
  
  export default reducer
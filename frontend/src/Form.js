import React, {Component} from 'react';
import './Form.css';

class Form extends Component {
  render() {
    return (
    <form>
      <h3>Форма авторизации</h3>
      <input type="text" placeholder="Login"></input>
      <input type="text" placeholder="Password"></input>
      <button type="submit">Confirm</button>
    </form>
  );
    }
}

// function Form () {
//       return (
//       <form>
//         <h3>Форма авторизации</h3>
//         <input type="text" placeholder="Login"></input>
//         <input type="text" placeholder="Password"></input>
//         <button type="submit">Confirm</button>
//       </form>
//     );
//   }
  

export default Form;
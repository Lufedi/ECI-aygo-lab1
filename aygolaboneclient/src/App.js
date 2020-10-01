import React, { useState } from 'react';
import './App.css';
import axios from 'axios'


const domain =  'http://localhost:5000/api/text';

const DisplayTexts = ({ texts }) => 
  <div>
    <ol>
      {
        texts.map((text,index) => <li
        style={{textAlign: "start"}}
        key={`${text}${index}`} >

           {`${text.Value}   created at:   ${new Date(text.CreatedDate).toDateString()} `}
        </li>)
      }
    </ol>
  </div>

const InputForm = ({ setTexts }) =>  {
  const [text, setText] = useState('')

  const sendText = async () => {
    const response = await axios.post( domain, { value : text })
    if(response.status == 200 && response.data){
      setTexts(response.data)
    }
  }
  return (
    <div>
      <input onChange={(e) =>{
        setText(e.target.value)
      } } /> 
      <button onClick={sendText}> send </button>
      <h6> Most recent texts: </h6>

    </div>
    
    
  )
}

const App = () => {

  const [texts, setTexts] = useState([])
  
  return (
    <div className="App">
      <header className="App-header">
        <p>
          Texts
        </p>
        <InputForm setTexts={setTexts}/>
        <DisplayTexts texts={texts} />        
      </header>
    </div>
  );
}

export default App;

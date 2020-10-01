import React, { useEffect, useState } from 'react';
import './App.css';
import axios from 'axios'


//const domain =  'http://ec2-34-209-136-168.us-west-2.compute.amazonaws.com:5000';
const domain = 'http://localhost:5000'
const router = `${domain}/api/text`

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
    const response = await axios.post( router, { value : text })
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
      <h6> Most recent logs: </h6>

    </div>
    
    
  )
}

const App = () => {

  const [texts, setTexts] = useState([])

    useEffect(() => {
      const getTexts = async () => {
        const response = await axios.get(router)
        if(response.status === 200 && response.data) {
          setTexts(response.data)
        }
      }
      getTexts()
    }, [])
  
  return (
    <div className="App">
      <header className="App-header">
        <p>
          Logs
        </p>
        <InputForm setTexts={setTexts}/>
        <DisplayTexts texts={texts} />        
      </header>
    </div>
  );
}

export default App;

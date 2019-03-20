import React, { Component } from 'react';
import './App.css';


class App extends Component {
  constructor(props) {
    super(props);
    {/*Données qui seront récupérées depuis le serveur*/}
    this.state = {commentaires: 0 ,moyenne:0 };
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">

        {/*A l'envoie du formulaire (entrée du lien et clic sur le bouton) l'URL sera envoyée au serveur GOLANG*/}
        <form onSubmit={this.function}>
        <label>
          Entrez le lien dont vous voulez connaitre les statistiques<br/>
        </label>
        <input id="lien"/>

        <br/><br/>

          <button className="App-button">
            Button
          </button>
          </form>

          <br/>

          {/*Affichage des données récupéréées*/}
          <div>
        Commentaires: {this.state.commentaires}<br/>
        Moyenne : {this.state.moyenne}/5
          </div>

      </header>
      </div>
    );
  }
}

export default App;

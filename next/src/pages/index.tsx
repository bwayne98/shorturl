import Head from 'next/head'
import { styled } from 'styled-components'


const MainContainer = styled.div`
  height: 100dvh;
  width: 100dvw;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
    
  > h3 {
      text-align: center;
      font-weight: 100;
      font-size: 2.6rem;
      padding: 15px 0px;
      letter-spacing: 2px;
      text-shadow: 1px 1px 5px gray;
      user-select: none;
    }
  > .inner{
    width: 450px;
    max-width: 100dvw;
    padding: 5px;
    text-shadow: 1px 1px 3px gray;


    > .input{
      width: 100%;
      display: grid;
      grid-template-columns: max-content 1fr;
      > label{
        padding: 0px 10px;
        font-size: 1.2rem;
        font-weight: 100;
        letter-spacing: 1px;
      }
      > input{
        outline: none;
        padding: 0 3px;
        border-radius: 1px;
        border-style: none;
        border-bottom-style: solid;
        border-color: linen;
        color: linen;
        background-color: transparent;
        letter-spacing: 1px;
      }
    }

    > .button{
      display: flex;
      justify-content:center;
      padding: 15px;
      > button{
        background-color: transparent;
        outline-style: none;
        padding: 3px 5px;
        border-width: 1px;
        outline-width: 0px;
        border-style: solid;
        border-radius: 2px;
        font-size: 1.1rem;
        font-weight: 100;
        letter-spacing: 1px;
        border-color: #c8c8c8;
        cursor: pointer;
        &:hover{
          background-color: rgba(200,200,200,.2);
        }
        &:active{
          opacity: .6;
        }
      }
    }
  }
`;

export default function Home() {
  return (
    <>
      <Head>
        <title>Zhmen Shorturl</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <MainContainer>
        <h3>Make A Short Url</h3>
        <div className="inner">
          <div className="input">
            <label htmlFor="">Origin Url</label>
            <input type="text" />
          </div>
          <div className="button">
            <button>Generate</button>
          </div>
        </div>
      </MainContainer>
    </>
  )
}

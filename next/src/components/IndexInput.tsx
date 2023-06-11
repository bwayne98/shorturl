"user client";

import { useRef, useState } from "react";
import { styled } from "styled-components";

export default function IndexInput() {

  const [result, setResult] = useState<string>('');
  const [clicked, setClicked] = useState<boolean>(false);
  const url = useRef<HTMLInputElement>(null);

  const onCLickGenerate = async () => {
    if (clicked) return;
    setClicked(true);

    const apiUrl = `https://${location.host}/api/short/make`;
    const data: ShortMakeBody = {
      origin: url.current?.value as string
    };
    const response = await fetchShortMake(apiUrl, data);
    if (response) {
      setResult(response.shortUrl);
    }

    setTimeout(() => {
      setClicked(false);
    }, 500);
  }

  const onClickCopy = () => {
    navigator.clipboard.writeText(result);
  }

  return (
    <Container clicked={clicked}>
      {result.length > 0 ?
        (
          <>
            <p>{result}</p>
            <div className="button">
              <button onClick={onClickCopy}>Copy</button>
              <button onClick={() => setResult("")}>Back</button>
            </div>
          </>
        ) : (
          <>
            <div className="input">
              <label htmlFor="">Origin Url</label>
              <input ref={url} type="text" placeholder="https://example.com" />
            </div >
            <div className="button">
              <button onClick={onCLickGenerate}>Generate</button>
            </div>
          </>
        )
      }
    </Container >
  )
}

const Container = styled.div<{ clicked: boolean }>`
    width: 450px;
    max-width: 100dvw;
    padding: 5px;
    text-shadow: 1px 1px 3px gray;
    overflow-wrap: break-word;

    > p {
      width: 100%;
      text-align: center;
      user-select: all;
      font-size: 1.2rem;
      max-width: 100%;
    }


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
        &::placeholder{
          color: #c8c8c8;
        }
      }
    }

    > .button{
      display: flex;
      justify-content:center;
      padding: 15px;
      gap: 15px;
      > button{
        background-color: ${props => props.clicked ? 'rgba(200,200,200,.2)' : 'transparent'};
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
        opacity: ${props => props.clicked ? '0.6' : '1'};
        pointer-events: ${props => props.clicked ? 'none' : 'auto'};
        &:hover{
          background-color: rgba(200,200,200,.2);
        }
        &:active{
          opacity: .6;
        }
      }
    }
`;

type ShortMakeResponseJson = {
  id: number,
  shortUrl: string,
  expiredAt: string,
}

type ShortMakeBody = {
  origin: string
}

const fetchShortMake = async (url: string, data: ShortMakeBody): Promise<ShortMakeResponseJson | false> => {
  try {
    const response = await fetch(url, {
      method: "POST",
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
    if (response.status === 200) {
      return await response.json();
    } else {
      throw await response.json();
    }
  } catch (e) {
    console.log(e)
    return false;
  }
}


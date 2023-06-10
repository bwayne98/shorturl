import Head from 'next/head'
import { styled } from 'styled-components'
import IndexInput from '../components/IndexInput';


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
        <IndexInput></IndexInput>
      </MainContainer>
    </>
  )
}

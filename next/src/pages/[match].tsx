import { GetServerSidePropsContext } from "next";

type Params = {
  match: string
}

export async function getServerSideProps({ params }: GetServerSidePropsContext) {
  const { match } = params as Params;
  try {
    const response = await fetch('http://go:8000/api/short/match', {
      method: 'POST',
      headers: {
        'content-type': 'application/json'
      },
      body: JSON.stringify({ match })
    })

    if (response.status !== 200) {
      throw Error('200');
    }

    const { redirectUrl } = await response.json();

    return {
      redirect: {
        destination: redirectUrl,
        permanent: false,
      }
    }

  } catch (e) {
    return {
      redirect: {
        destination: '/',
        permanent: false,
      }
    }
  }
}

type Props = {
  match: string
}


export default function Match({ match }: Props) {
  return (
    <>
      <h1>{JSON.stringify(match)}</h1>
    </>
  )
}
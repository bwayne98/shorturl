import App, { AppProps } from "next/app";
import React, { ReactElement, ReactNode } from "react";
import { NextPage } from "next";

import { parseCookies } from "nookies";
import '../styles/globals.css';

// 為 type NextPage 增加 getLayout 屬性
type NextPageWithLayout<P = {}, IP = P> = NextPage<P, IP> & {
  getLayout?: (page: ReactElement) => ReactNode;
};

// AppProps 為 App預設props接收的型別
// 複寫 AppProps.component 型別 (  )
// 同時擴加 cookies
type AppPropsWithLayout = AppProps & {
  Component: NextPageWithLayout;
  cookies: any;
};

export default function MyApp({
  Component,
  pageProps,
  cookies,
}: AppPropsWithLayout) {

  // 判斷component 是否有 getLayout
  // 沒有就使用預設的 Layout
  // 等於提供 Page 複寫 Layout 的路口
  const PageLayout = Component.getLayout ?? ((page) => page);

  return PageLayout(<Component {...pageProps} />);
}

// 自定義 App ,提取 appProps ,cookies
MyApp.getInitialProps = async (appContext: any) => {
  const appProps = await App.getInitialProps(appContext);
  const cookies = parseCookies(appContext.ctx);

  return { ...appProps, cookies };
};
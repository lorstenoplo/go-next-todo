import { TextField } from "@mui/material";
import Head from "next/head";
import { Header } from "../components/Header";

export default function Home() {
  return (
    <div style={{ backgroundColor: "#121212", height: "100vh" }}>
      <Head>
        <title>Todo app</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Header />

      <TextField id="outlined-basic" label="Outlined" variant="outlined" />
    </div>
  );
}

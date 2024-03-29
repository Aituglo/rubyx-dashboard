import React, { useState, useEffect } from "react";
import PageTitle from "../Typography/PageTitle";
import {
  Table,
  TableHeader,
  TableCell,
  TableBody,
  TableRow,
  TableFooter,
  TableContainer,
  Button,
  Pagination,
} from "@windmill/react-ui";
import ClipLoader from "react-spinners/ClipLoader";
import Input from "../Input";
import { getUrls } from "../../actions/url";
import { useSelector, useDispatch } from "react-redux";
import { TbArrowBack } from "react-icons/tb";

function Url({ setSeeUrls, url }) {
  const dispatch = useDispatch();
  const urlState = useSelector((state) => state.url);
  const [searchTerm, setSearchTerm] = useState("");

  const resultsPerPage = 50;

  useEffect(() => {
    dispatch(getUrls(url, 1, resultsPerPage, searchTerm));
  }, [url, searchTerm]);

  const onPageChangeTable = (p) => {
    dispatch(getUrls(url, p, resultsPerPage, searchTerm));
  };


  return (
    <>
      <PageTitle>Urls</PageTitle>

      <div className="px-4 py-3 mb-8 bg-white rounded-lg shadow-md dark:bg-gray-800">
        <div className="flex items-center mb-4 space-x-4">
          <Button className="bg-blue-900" onClick={() => setSeeUrls(false)}>
            <TbArrowBack className="w-5 h-5" />
          </Button>

          <Input
            className="text-gray-700"
            placeholder="Search for url"
            aria-label="Search"
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>

        {urlState.loading && (
          <div className="flex items-center justify-center">
            <ClipLoader
              color={"#355dad"}
              loading={true}
              size={30}
              aria-label="Loading"
              data-testid="loader"
            />
          </div>
        )}

        {urlState.total && urlState.total > 0 ? (
          <TableContainer className={`mb-8 ${urlState.loading && "hidden"}`}>
            <Table>
              <TableHeader>
                <tr>
                  <TableCell>Url</TableCell>
                  <TableCell>Status Code</TableCell>
                </tr>
              </TableHeader>
              <TableBody>
                {urlState.urls &&
                  urlState.urls.map((key, i) => (
                    <TableRow key={i}>
                      <TableCell>
                        <span className="text-sm">
                          <a
                            href={key.url}
                            target="_blank"
                            rel="noreferrer noopener"
                          >
                            {key.url}
                          </a>
                        </span>
                      </TableCell>
                      <TableCell>
                        <span className="text-sm">{key.status_code}</span>
                      </TableCell>
                    </TableRow>
                  ))}
              </TableBody>
            </Table>
            <TableFooter>
              <Pagination
                totalResults={urlState.total ? urlState.total : 0}
                resultsPerPage={resultsPerPage}
                onChange={onPageChangeTable}
                label="Navigation"
              />
            </TableFooter>
          </TableContainer>
        ) : (
          <div className="flex items-center justify-center">
            <span className="text-sm">No data to display</span>
          </div>
        )}
      </div>
    </>
  );
}

export default Url;
